var days = 256;
var fish_array = [1,5,5,1,5,1,5,3,1,3,2,4,3,4,1,1,3,5,4,4,2,1,2,1,2,1,2,1,5,2,1,5,1,2,2,1,5,5,5,1,1,1,5,1,3,4,5,1,2,2,5,5,3,4,5,4,4,1,4,5,3,4,4,5,2,4,2,2,1,3,4,3,2,3,4,1,4,4,4,5,1,3,4,2,5,4,5,3,1,4,1,1,1,2,4,2,1,5,1,4,5,3,3,4,1,1,4,3,4,1,1,1,5,4,3,5,2,4,1,1,2,3,2,4,4,3,3,5,3,1,4,5,5,4,3,3,5,1,5,3,5,2,5,1,5,5,2,3,3,1,1,2,2,4,3,1,5,1,1,3,1,4,1,2,3,5,5,1,2,3,4,3,4,1,1,5,5,3,3,4,5,1,1,4,1,4,1,3,5,5,1,4,3,1,3,5,5,5,5,5,2,2,1,2,4,1,5,3,3,5,4,5,4,1,5,1,5,1,2,5,4,5,5,3,2,2,2,5,4,4,3,3,1,4,1,2,3,1,5,4,5,3,4,1,1,2,2,1,2,5,1,1,1,5,4,5,2,1,4,4,1,1,3,3,1,3,2,1,5,2,3,4,5,3,5,4,3,1,3,5,5,5,5,2,1,1,4,2,5,1,5,1,3,4,3,5,5,1,4,3];
var sum = 0;
var map = new Map();
var fish_counters = new Array(9);
fish_counters.fill(0);

/*for (let i = 0; i < fish_array.length; i++) {
	if (map.has(fish_array[i])) {
		let val = map.get(fish_array[i]);
		map.set(fish_array[i], val + 1);
	}
	else {
		map.set(fish_array[i], 1);
	}
}*/

for (let i = 0; i < fish_array.length; i++) {
	fish_counters[fish_array[i]]++;
}
sum += fish_array.length;

console.log(fish_counters);

for (let day = 0; day < days; day++) {
	const new_fish = fish_counters[0];
	fish_counters.shift();
	fish_counters.push(0);
	fish_counters[6] += new_fish;
	fish_counters[8] += new_fish;
	sum += new_fish;
}

/*for (const [starting_pos, quantity] of Object.entries(map)) {
	const fish_count = CalculateFish(starting_pos, days);

	sum += fish_count * quantity;
}

function CalculateFish(starting_pos, days) {
	days--;


}*/

/*var cycles = Math.ceil((days - array[i]) / 6);
	sum += cycles;

	for (let j = 0; j < cycles + 1; j++) {
		let new_days = days - array[i] - j * 6;
		//console.log("j " + j + " " + new_days);
		sum += CalculateChild(new_days);
	}*/

/*function CalculateFish(starting_pos) {
	//console.log("day: " + new_days);
	if (new_days <= 0) {
		console.log("return 0");
		return 0;
	}

	if (starting_pos > 0) {

	}

	var new_sum = 0;

	if (new_days - 8 >= 0) {
		new_sum += 1;

		cycles = Math.ceil((new_days - 8) / 6);

			for (let i = 0; i < cycles + 1; i++) {
			console.log("for: " + i)
			new_sum += CalculateChild(new_days - 8);
		}
	}

	console.log(new_days);
	console.log("return end " + new_sum);

	return new_sum;
}*/

console.log(sum);
