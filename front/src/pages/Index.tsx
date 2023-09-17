import { Box } from '@mui/material';
import { Link } from '../components/Link';

export function Index() {
  return (
    <Box
      sx={{
        width: '100%',
        height: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <h1>
        <Link
          to="/pick-user"
          style={{
            fontFamily: 'Allerta Stencil',
            fontSize: 55,
          }}
        >
          BOTTLIS
        </Link>
      </h1>
    </Box>
  );
}
