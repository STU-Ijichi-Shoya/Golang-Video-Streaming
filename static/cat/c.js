'use strict';

//require
const Obniz = require('obniz');
const request = require('request');
const { createCanvas } = require('canvas');

//LINE Notify URL
const LINE_NOTIFY_URL = 'https://notify-api.line.me/api/notify';
//LINE Notify トークン
const TOKEN = 'kG06LbcBh9FLDwVuCJZ1IznURKy81M47W8thgHeXL9j';

//時間取得 月、日、時、分、秒
function get_time() {
  var now = new Date();
  var Month = now.getMonth() + 1;
  var Dates = now.getDate();
  var Hour = now.getHours();
  var Min = now.getMinutes();
  var Sec = now.getSeconds();

  var times = Month + "月" + Dates + "日" + Hour + ":" + Min + ":" + Sec;
  return times;
};

function getOpenMsg(){
    //開いたときのMESSAGE
    return '部室が開きました\n' + get_time() + "\n" + "換気扇をつけて、窓を開けて換気をしましょう！";
}
//閉じたときのMESSAGE
function getCloseMsg(){
    return '部室が閉まりました\n' + get_time() + "\n" + "換気扇を消して、窓を閉めましょう！";
}
function getOpenOption(){
    //開いたときのOPTION
    let m=getOpenMsg()
    return {
    url: LINE_NOTIFY_URL,
    method: 'POST',
    headers: HEADERS,
    json: true,
    form: { message: m }
  }
}
var HEADERS = {
  'Content-Type': 'application/x-www-form-urlencoded',
  'Authorization': 'Bearer ' + TOKEN
};

function getCloseOption(){
    let m=getCloseMsg()
    return {
        url: LINE_NOTIFY_URL,
        method: 'POST',
        headers: HEADERS,
        json: true,
        form: { message: m }
      }      
}

//ObnizID
var obniz = new Obniz("1325-2702");

//オンライン時
obniz.onconnect = async function () {
  //ディスプレイクリア
  obniz.display.clear();
  const canvas = createCanvas(128, 64);
  const ctx = canvas.getContext('2d');
  ctx.fillStyle = "white";
  ctx.font = "40px Avenir";

  //LED 0ピン1ピン
  var led = obniz.wired("LED", { anode: 0, cathode: 1 });
  //ドア磁気センサー2ピン3ピン
  var button = obniz.wired("Button", { signal: 2, gnd: 3 });

  //閉じたとき
  button.onchange = function (pressed) {
    if (pressed) {
      console.error(get_time());
      request(getCloseOption(), (error, request, body) => {
        if (error) {
          console.error(error);
          return;
        }
        console.log(body);
        ctx.clearRect(0, 0, 128, 64)
        //ディスプレイにCLOSE出力
        ctx.fillText('CLOSE', 0, 40);
        obniz.display.draw(ctx);
        //LED OFF
        led.off();
      });

    }

    //閉じたとき
    else {
      console.error(get_time());
      request(getOpenOption(), (error, request, body) => {
        if (error) {
          console.error(error);
          return;
        }
        console.log(body);
        ctx.clearRect(0, 0, 128, 64)
        //ディスプレイにOPEN出力
        ctx.fillText('OPEN', 0, 40);
        obniz.display.draw(ctx);
        //LED ON
        led.on();
      });

    };
  };
};

//オフライン時
obniz.onclose = async function () {

};