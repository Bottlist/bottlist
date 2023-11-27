import './index.css';
import { Index } from './pages/Index';
import { PickUser } from './pages/PickUser';
import { bottleRouter } from './pages/bottles/router';
import { shopsRouter } from './pages/shops/router';
import { usersRouter } from './pages/users/router';
import reportWebVitals from './reportWebVitals';
import * as serviceWorkerRegistration from './serviceWorkerRegistration';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const router = createBrowserRouter([
  {
    path: '/',
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
        path: 'users',
        children: usersRouter,
      },
      {
        path: 'shops',
        children: shopsRouter,
      },
      {
        path: 'bottle',
        children: bottleRouter,
      },
    ],
  },
]);

const theme = createTheme({
  typography: {
    fontFamily: 'Yu Gothic Medium',
  },
  components: {
    MuiTextField: {
      defaultProps: { sx: { width: '100%' } },
    },
    MuiButton: {
      defaultProps: {
        variant: 'contained',
      },
    },
  },
  palette: {
    background: {
      default: '#EAE8E1',
    },
    primary: {
      main: '#D27C2C',
    },
    secondary: {
      main: '#EAE8E1',
    },
  },
});

const queryClient = new QueryClient();

root.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <QueryClientProvider client={queryClient}>
          <CssBaseline />
          <RouterProvider router={router} />
        </QueryClientProvider>
      </LocalizationProvider>
    </ThemeProvider>
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
