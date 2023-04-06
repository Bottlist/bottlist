import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';
import { Outlet } from 'react-router-dom';

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
    primary: {
      main: '#D0EF9D',
    },
    info: {
      main: '#00FF00',
    },
  },
});

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Outlet />
    </ThemeProvider>
  );
}
export default App;
