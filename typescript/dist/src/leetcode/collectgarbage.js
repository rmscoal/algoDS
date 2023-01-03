/* eslint-disable guard-for-in */
/* eslint-disable no-restricted-syntax */
/* eslint-disable no-param-reassign */
/* eslint-disable no-plusplus */
/* eslint-disable default-case */
export default class CollectGarbage {
    /**
     * @public @static
     * @param garbage garbage in each house
     * @param travel time takes to travel to each houses
     */
    static Solver(garbage, travel) {
        const garbageData = {
            glass: {
                letter: 'G',
                maxHouseFound: false,
                maxHouseIndex: -1,
                min: 0,
            },
            paper: {
                letter: 'P',
                maxHouseFound: false,
                maxHouseIndex: -1,
                min: 0,
            },
            metal: {
                letter: 'M',
                maxHouseFound: false,
                maxHouseIndex: -1,
                min: 0,
            },
        };
        // Finds the maximum index of a house each
        // garbage collector has to look for
        for (let i = garbage.length - 1; i >= 0; i -= 1) {
            if (garbage[i].includes('G')) {
                if (!garbageData.glass.maxHouseFound) {
                    garbageData.glass.maxHouseFound = true;
                    garbageData.glass.maxHouseIndex = i;
                }
            }
            if (garbage[i].includes('P')) {
                if (!garbageData.paper.maxHouseFound) {
                    garbageData.paper.maxHouseFound = true;
                    garbageData.paper.maxHouseIndex = i;
                }
            }
            if (garbage[i].includes('M')) {
                if (!garbageData.metal.maxHouseFound) {
                    garbageData.metal.maxHouseFound = true;
                    garbageData.metal.maxHouseIndex = i;
                }
            }
        }
        // Returned variable:
        let min = 0;
        // Do a for loop for each of the garbage type
        // to calculates the minimum time based on
        // the maximum index of the house
        for (const key in garbageData) {
            for (let i = 0; i <= garbageData[key].maxHouseIndex; i += 1) {
                if (garbage[i].includes(garbageData[key].letter)) {
                    garbageData[key].min += (garbage[i].match(new RegExp(`${garbageData[key].letter}`, 'g'))
                        || []).length;
                }
                if (i < garbageData[key].maxHouseIndex) {
                    garbageData[key].min += travel[i];
                }
            }
            min += garbageData[key].min;
        }
        return min;
    }
}
//# sourceMappingURL=collectgarbage.js.map