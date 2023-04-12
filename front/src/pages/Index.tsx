import { VerticalCenter } from '../components/VerticalCenter';
import { Typography } from '@mui/material';
import { useNavigate } from 'react-router';

export function Index() {
  const navigate = useNavigate();
  return (
    <VerticalCenter onClick={() => navigate('/pick-user')}>
      <Typography variant="h1" textAlign="center">
        Bottlist
      </Typography>
    </VerticalCenter>
  );
}
