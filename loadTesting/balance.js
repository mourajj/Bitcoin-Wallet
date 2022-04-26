import http from 'k6/http';
import { sleep } from 'k6';


export let options ={
    noConnectionReuse:false,
    vus: 1, //An integer value specifying the number of VUs (Virtual users) to run concurrently, used together with the iterations or duration options.
    duration: '1s' //Duration of the test
};

export default function (){

   // http.get("https://api.coinbase.com/v2/prices/USD-EUR/spot") If you want to test the delay of the external API

    http.get('http://127.0.0.1:5000/balance/pedro') //Its recommended to specify an user that is present in the database, otherwise the call will not be properly made.
    
    //sleep(1)
}
