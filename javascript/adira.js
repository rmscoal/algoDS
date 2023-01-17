const https = require('https');

https.get('https://coderbyte.com/api/challenges/json/json-cleaning', (resp) => {

  let data = '';

  // parse json data here...
  // Checker on == '' || == 'N\/A' || == '-' could be improved into its
  // own function to maintain clean code principal.
  resp.on("data", (body) => {
    let respBody = JSON.parse(body);
    let keys = Object.keys(respBody)
    for (let i = 0; i < keys.length; i++) {
      if (typeof respBody[keys[i]] == "object") {
        if (Array.isArray(respBody[keys[i]])) {
          for (let j = 0; j < respBody[keys[i]].length; j++) {
            if (respBody[keys[i]][j] == '' || respBody[keys[i]][j] == 'N\/A' || respBody[keys[i]][j] == '-') {
              respBody[keys[i]].splice(j, 1)
            }
          }
        } else {
          let childKeys = Object.keys(respBody[keys[i]])
          for (let j = 0; j < childKeys.length; j++) {
            if (respBody[keys[i]][childKeys[j]] == '' || respBody[keys[i]][childKeys[j]] == '-' || respBody[keys[i]][childKeys[j]] == 'N\/A') {
              delete respBody[keys[i]][childKeys[j]]
            }
          }
        }
      } else {
        if (respBody[keys[i]] == '' || respBody[keys[i]] == 'N\/A' || respBody[keys[i]] == '-') {
          delete respBody[keys[i]]
        }
      }
    }

    data += JSON.stringify(respBody)
    console.log(data)
  })

  resp.on("error", (err) => {
    console.log("Err " + err)
  })
});