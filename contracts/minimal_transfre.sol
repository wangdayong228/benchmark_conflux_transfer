pragma solidity ^0.8.0;

contract SimpleTransfer {
    event Transfer(address indexed from, address indexed to, uint256 value);

    function run(
        address from,
        address to,
        uint256 repeatCount
    ) public {
        for (uint256 i = 0; i < repeatCount; i++) {
            emit Transfer(from, to, i);
        }
    }
}
