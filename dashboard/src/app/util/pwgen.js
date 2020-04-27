const jsSHA = require("jssha");
export function genPW() {
  let key;
  let isValid = false;
  while (!isValid) {
    key = keyGen();
    const firstFive = key["hashKey"].slice(0, 5);
    const getUrl = "https://api.pwnedpasswords.com/range/" + firstFive;
    const data = getData(getUrl);
    const jsonData = handleData(data);
    isValid = checkForMatch(jsonData, key);
  }
  return key["clearKey"];
}

function shuffle(array) {
  let currentIndex = array.length, temporaryValue, randomIndex;
  while (0 !== currentIndex) {
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }
  return array;
}

function keyGen() {
  let chars = [];
  let i;
  for (i = 48; i < 58; i++) {chars.push(String.fromCharCode(i));}
  for (i = 65; i < 91; i++) {chars.push(String.fromCharCode(i));}
  for (i = 97; i < 122; i++) {chars.push(String.fromCharCode(i));}
  chars.push(String.fromCharCode(33));
  chars.push(String.fromCharCode(95));
  chars = shuffle(chars);
  let key = "";
  for (i = 0; i < 24; i++) {
    var rannumber = Math.floor(Math.random() * chars.length);
    key += chars[rannumber]
  }
  const shaObj = new jsSHA("SHA-1", "TEXT");
  shaObj.update(key);
  let hashKey = shaObj.getHash("HEX");
  return {"clearKey":key,"hashKey":hashKey};
}

function getData(getUrl) {
  let jsonData = null;
  $.ajax({
    type: "GET",
    async: false,
    url: getUrl,
    success: function (data) {
      jsonData = data;
    }
  });
  return jsonData;
}

function handleData(data) {
  let splitData = data.split("\n");
  let jsonData = [];
  let i;
  for (i = 0; i < splitData.length; i++) {
    let hash = splitData[i].slice(0,35);
    let count = splitData[i].slice(36);
    jsonData.push({hash : hash, count : count});
  }
  return jsonData;
}

function checkForMatch(jsonData, key) {
  let checkKey = key["hashKey"].slice(5);
  let retValue = false;
  let i;
  for (i = 0; i < jsonData.length; i++) {
    let obj = jsonData[i];
    if (obj.hash === checkKey.toUpperCase()) {
      retValue = false;
      console.log("MATCH FOUND AT " + i + "\n occurred " + obj.count + " times");
    } else {
      retValue = true;
    }
  }
  return retValue;
}
