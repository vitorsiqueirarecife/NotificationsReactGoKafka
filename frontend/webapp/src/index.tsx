import React from 'react';
import ReactDOM from 'react-dom/client';
import Notification from './module/Notification/screen/index';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <Notification />
  </React.StrictMode>
);

reportWebVitals();
