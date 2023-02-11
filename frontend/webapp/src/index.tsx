import React from 'react';
import { QueryClientProvider, QueryClient } from '@tanstack/react-query';
import ReactDOM from 'react-dom/client';
import NotificationScreen from './module/Notification/screen/index';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const queryClient = new QueryClient()

root.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <NotificationScreen />
    </QueryClientProvider>
  </React.StrictMode>
);

reportWebVitals();
