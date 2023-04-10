import { Typography } from '@mui/material';
import { useNavigate } from 'react-router';
import { VerticalCenter } from '../components/VerticalCenter';

export function Index() {
  const navigate = useNavigate();
  return (
    <VerticalCenter onClick={() => navigate('/pick-user')}>
      <Typography variant="h1">Bottlist</Typography>
    </VerticalCenter>
  );
}
