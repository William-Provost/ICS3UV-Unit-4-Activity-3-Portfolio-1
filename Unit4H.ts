/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program asks the user for their 10 best friends and displays them in order and reverse order.
 */

// variables
const bestFriends: string[] = new Array(10);

console.log("Please enter your 10 best friends in order from oldest friend to newest friend:\n");

// get 10 friends from user
for (let i = 0; i < 10; i++) {
  bestFriends[i] = prompt(`Enter friend ${i + 1}: `) || "";
}

// display friends in order
console.log("\nHere is your list of best friends in order of length of friendship starting from oldest friend to newest friend:");
console.log(bestFriends.join(", ") + ".");

// display friends in reverse order
console.log("\nHere is your list of best friends in reverse order of length of friendship starting from newest friend to oldest friend:");
console.log(bestFriends.slice().reverse().join(", ") + ".");

console.log("\nDone.");
