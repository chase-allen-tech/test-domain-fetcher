import React, { useCallback, useEffect, useState } from "react";
import axios from "axios";
import "./App.css";

type Status = {
  url: string;
  statusCode: number;
  duration: number;
  date: string;
};

function App() {
  const [amazonStats, setAmazonStatus] = useState<Status>();
  const [googleStats, setGoogleStatus] = useState<Status>();
  const [allStats, setAllStatus] = useState<Status[]>([]);

  const doFetch = useCallback(() => {
    axios
      .get(`${process.env.REACT_APP_API_ENDPOINT}/amazon-status`)
      .then((res) => setAmazonStatus(res.data))
      .catch((err) => console.error(err));

    axios
      .get(`${process.env.REACT_APP_API_ENDPOINT}/google-status`)
      .then((res) => setGoogleStatus(res.data))
      .catch((err) => console.error(err));

    axios
      .get(`${process.env.REACT_APP_API_ENDPOINT}/all-status`)
      .then((res) => setAllStatus(res.data))
      .catch((err) => console.error(err));

    setTimeout(doFetch, 60*1000)
  }, []);

  useEffect(() => {
    doFetch();
  }, [doFetch]);

  return (
    <div className="App">
      <div>
        <h3>Amazon</h3>
        <h5>{amazonStats?.url}</h5>
        <h5>{amazonStats?.statusCode}</h5>
        <h5>{amazonStats?.duration} ms</h5>
        <h5>{amazonStats?.date}</h5>
      </div>
      <div>
        <h3>Google</h3>
        <h5>{googleStats?.url}</h5>
        <h5>{googleStats?.statusCode}</h5>
        <h5>{googleStats?.duration} ms</h5>
        <h5>{googleStats?.date}</h5>
      </div>
      <div>
        <h3>All</h3>
        {allStats.map((v) => (
          <div key={v.url}>
            <h5>{v.url}</h5>
            <h5>{v.statusCode}</h5>
            <h5>{v.duration} ms</h5>
            <h5>{v.date}</h5>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
