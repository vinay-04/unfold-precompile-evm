// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface IZkSNARK {
  // verifyProof verifies the provided zkSNARK proof
  function verifyProof(
    uint256[] calldata input,
    uint256[] calldata proof,
    uint256[] calldata publicSignals
  ) external view returns (bool success);

  // generateProof generates a zkSNARK proof for the given inputs
  function generateProof(uint256[] calldata input) external returns (uint256[] memory proof);
}