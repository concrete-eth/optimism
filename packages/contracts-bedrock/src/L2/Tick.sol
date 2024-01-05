// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

interface ITick {
    function tick() external;
}

/// @custom:proxied
/// @custom:predeploy 0x42000000000000000000000000000000000000f0
/// @title Tick
/// @notice The Tick predeploy ticks the chain.
contract Tick is Ownable, ITick, ISemver {
    /// @custom:semver 0.1.0
    string public constant version = "0.1.0";

    /// @notice Address of the special depositor account.
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    /// @notice Address of the tick contract to be called.
    address public target;

    /// @notice Allows the owner to modify the target address.
    /// @param _target New target address.
    function setTarget(address _target) public onlyOwner {
        target = _target;
    }

    /// @notice Calls the tick function in the target contract.
    function tick() external {
        require(msg.sender == DEPOSITOR_ACCOUNT, "Tick: only the depositor account can tick");
        if (target == address(0)) {
            return;
        }
        (bool success, ) = target.call(abi.encodeWithSignature("tick()"));
        require(success);
    }
}
