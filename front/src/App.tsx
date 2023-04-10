import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';
import { Outlet } from 'react-router-dom';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';

const theme = createTheme({
  components: {
    MuiTypography: {
      defaultProps: {
        align: 'center',
        justifySelf: 'center',
      },
    },
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

function App() {
  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <CssBaseline />
        <Outlet />
      </LocalizationProvider>
    </ThemeProvider>
  );
}
export default App;
