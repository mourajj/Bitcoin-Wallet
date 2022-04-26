import http from 'k6/http';
import { sleep } from 'k6';


export let options ={
    noConnectionReuse:false,
    vus: 1, //An integer value specifying the number of VUs (Virtual users) to run concurrently, used together with the iterations or duration options.
    duration: '5s'
};

export default function (){

    http.get('http://127.0.0.1:5000/balance/jonathan')
    //sleep(1)
}