function readAndParseFile(fileName) {
  var fs = require("fs"),
    path = require("path"),
    filePath = path.join(__dirname, fileName);
  return fs.readFileSync(filePath).toString().split("\n");
}

function mainFunction() {
  const array = readAndParseFile("input.txt");
  var highestCalories = [0,0,0];
  var currentElfCaolories = 0;
  array.forEach((calorie) => {
    if (calorie) {
      currentElfCaolories += parseInt(calorie);
    } else {
      if(currentElfCaolories>highestCalories[0]){
        highestCalories[0]=currentElfCaolories;
        highestCalories.sort((a,b)=>(a-b))
      }
      currentElfCaolories=0
    }
  });
  console.log(highestCalories.reduce((sum,currenetValue)=>{return sum+currenetValue}));
}

mainFunction();
