import { Assert } from "./test";

export function sumProdDiags(matrix: number[][]): number {
  const dimension = matrix.length;
  let products_sum = 0;

  matrix.map((_, offset) => {
    let p1 = 1;
    let m1 = 1;
    let p2 = Number(!!offset);
    let m2 = Number(!!offset); // So that products of first diagonals are not duplicated

    Array.from({ length: dimension - offset }).map((_, j) => {
      // Major diagonals
      p1 *= matrix[j][j + offset];
      p2 *= matrix[j + offset][j];
      // Minor diagonals
      m1 *= matrix[dimension - j - 1][j + offset];
      m2 *= matrix[dimension - j - 1 - offset][j];
    });
    products_sum += p1 + p2 - m1 - m2;
  });
  return products_sum;
}

const M1 = [
  [1, 4, 7, 6, 5],
  [-3, 2, 8, 1, 3],
  [6, 2, 9, 7, -4],
  [1, -2, 4, -2, 6],
  [3, 2, 2, -4, 7],
];
Assert("Running test with M1", sumProdDiags(M1), 1098);

const M2 = [
  [1, 4, 7, 6],
  [-3, 2, 8, 1],
  [6, 2, 9, 7],
  [1, -2, 4, -2],
];
Assert("Running test with M2", sumProdDiags(M2), -11);

const M3 = [
  [1, 2, 3, 2, 1],
  [2, 3, 4, 3, 2],
  [3, 4, 5, 4, 3],
  [4, 5, 6, 5, 4],
  [5, 6, 7, 6, 5],
];
Assert("Running test with M3", sumProdDiags(M3), 0);
