/**
 * A class to solve Number of Laser Beams in a Bank, leetcode, strings
 * Problem:
 * Anti-theft security devices are activated inside a bank. You are given a 0-indexed binary
 * string array bank representing the floor plan of the bank, which is an m x n 2D matrix.
 * bank[i] represents the ith row, consisting of '0's and '1's. '0' means the cell is empty,
 * while'1' means the cell has a security device.
 * There is one laser beam between any two security devices if both conditions are met:
 * 1. The two devices are located on two different rows: r1 and r2, where r1 < r2.
 * 2. For each row i where r1 < i < r2, there are no security devices in the ith row.
 * 3. Laser beams are independent, i.e., one beam does not interfere nor join with another.

Return the total number of laser beams in the bank.
 */
export default class MaxLaserBeam {
    static Solver(bank) {
        // The total of the laser beams
        let total = 0;
        // The first number of laser beams
        let prevNoOfDevices = 0;
        // In the first for loop, we are going to find the number of laser beams in the first row.
        for (let j = 0; j < bank[0].length; j += 1) {
            if (bank[0][j] === '1') {
                prevNoOfDevices += 1;
            }
        }
        // console.debug(`The first number of devices is ${prevNoOfDevices}`);
        for (let i = 1; i < bank.length; i += 1) {
            if (bank[i].includes('1')) {
                // console.debug(`Row ${i} contains 1`);
                let currentNoOfDevices = 0;
                for (let j = 0; j < bank[i].length; j += 1) {
                    if (bank[i][j] === '1') {
                        currentNoOfDevices += 1;
                    }
                }
                // console.debug(`Current No of Devices in this row is ${currentNoOfDevices}`);
                total += currentNoOfDevices * prevNoOfDevices;
                prevNoOfDevices = currentNoOfDevices;
            }
        }
        return total;
    }
}
//# sourceMappingURL=maxlaserbeam.js.map