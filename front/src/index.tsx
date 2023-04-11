import App from './App';
import './index.css';
import { Index } from './pages/Index';
import { Login } from './pages/Login';
import { PickLogin } from './pages/PickLogin';
import { PickUser } from './pages/PickUser';
import { Register } from './pages/Register';
import { Top } from './pages/Top';
import { bottleRouter } from './pages/bottles/router';
import reportWebVitals from './reportWebVitals';
import * as serviceWorkerRegistration from './serviceWorkerRegistration';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      {
        path: '',
        element: <Index />,
      },
      {
        path: 'pick-user',
        element: <PickUser />,
      },
      {
        path: 'pick-login',
        element: <PickLogin />,
      },
      {
        path: 'login',
        element: <Login />,
      },
      {
        path: 'register',
        element: <Register />,
      },
      {
        path: 'top',
        element: <Top />,
      },
      {
        path: 'bottle',
        children: bottleRouter,
      },
    ],
  },
]);

root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorkerRegistration.register();
