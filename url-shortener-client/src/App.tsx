import { useState } from 'react';
import './App.css'
import { BiLink } from 'react-icons/bi';
import { BsClipboard } from 'react-icons/bs';
import axios, { AxiosRequestConfig, AxiosRequestHeaders } from 'axios';
import { ShortenRequest, ShortenResponse } from './models/Models';

const App = () => {
  const [urlToShort, setUrl] = useState<string>('');
  const [shortUrl, setShortUrl] = useState<string>('');

  const copyToClipboard = () => {
    navigator.clipboard.writeText(shortUrl);
  }

  const handleClick = async () => {
    try {
      const reqBody: ShortenRequest = {
        url: urlToShort
      };
      const headers = {
        headers: {
          'Content-Type': 'application/json'
        } as AxiosRequestHeaders
      } as AxiosRequestConfig;
      const { data } = await axios.post<ShortenResponse>('http://localhost:8080/shorten', reqBody, headers);
      const { short_url } = data as ShortenResponse;

      setShortUrl(short_url);
    } catch (err) {
      if (axios.isAxiosError(err))
        console.error(err.message);
      else
        console.error(`Unexpected error occurred: ${err}`)
    }
  };

  return (
    <div id="container">
      <div id="bgImg"></div>
      <div id="urlBox" className='border'>
        <div id="header">
          <BiLink id='linkIcon' />
          <h1 id='title'>Url Shortener</h1>
        </div>
        <div id='toShorten'>
          <label htmlFor="urlToShorten">Url to Shorten</label>
          <input
            type="text"
            name="urlToShorten"
            id="urlToShorten"
            className='border input'
            onChange={(e) => setUrl(e.target.value)}
          />
        </div>

        <div id='shortened'>
          <label htmlFor="shortenedUrl">Shortened Url</label>
          <div id="shortenedLabelContainer">
            <button id='clipboardButton' onClick={copyToClipboard}>
              <BsClipboard id='clipboardIcon' />
            </button>
            <input
              type="text"
              name="shortenedUrl"
              id='shortenedUrl'
              className='border input'
              value={shortUrl}
              readOnly
            />
          </div>
        </div>

        <button onClick={handleClick}>Shorten Url</button>
      </div>
    </div>
  );
};

export default App;