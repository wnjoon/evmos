// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.0;

import "./IPacketActor.sol";

contract PacketActorCounter is PacketActorBase  {

    int public counter = 0;

    function onSendPacket(
        ICS20Packet calldata packet,
        address relayer
    ) external returns (bool success) {
        counter += 1;
        return true;
    }

    function onRecvPacket(
        ICS20Packet calldata packet
    ) external returns (bool success) {
        return true;
    }

    function onAcknowledgementPacket(
        ICS20Packet calldata packet,
        bytes calldata acknowledgement,
        address relayer
    ) external returns (bool success) {
        counter += 1;
        return true;
    }

    function onTimeoutPacket(
        ICS20Packet calldata packet,
        address relayer
    ) external returns (bool success) {
        return true;
    }
}