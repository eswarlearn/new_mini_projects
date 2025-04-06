import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { Passingdatas } from "./Components/DataContext";

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Passingdatas>
      <App />
    </Passingdatas>
  </React.StrictMode>
);

reportWebVitals();
