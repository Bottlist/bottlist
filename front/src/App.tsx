import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Outlet } from 'react-router-dom';

const theme = createTheme({
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
      main: '#D9D9D9',
    },
    secondary: {
      main: '#FFFFFF',
    },
  },
});

const queryClient = new QueryClient();

function App() {
  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <QueryClientProvider client={queryClient}>
          <CssBaseline />
          <Outlet />
        </QueryClientProvider>
      </LocalizationProvider>
    </ThemeProvider>
  );
}
export default App;
