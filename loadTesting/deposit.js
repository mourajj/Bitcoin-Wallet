import http from 'k6/http';
import { sleep } from 'k6';

export let options ={
    noConnectionReuse:false,
    vus: 1, //An integer value specifying the number of VUs (Virtual users) to run concurrently, used together with the iterations or duration options.
    duration: '60s'
};

export default function (){

    const payload = JSON.stringify({
        "user":"testing",
        "currency":"btc",
        "amount":1
    })

    var params = {
        headers:{
            'Content-type':'application/json'
        }
    }

    http.post('http://127.0.0.1:5000/deposit',payload,params)
    //sleep(1)

}