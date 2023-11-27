import { Container, Modal as MuiModal, Paper } from '@mui/material';

export const Modal = (props: {
  children: JSX.Element | JSX.Element[];
  open: boolean;
  onClose: () => void;
  width?: string;
}) => (
  <MuiModal
    open={props.open}
    onClose={props.onClose}
    sx={{ top: '50%', width: props.width, marginX: 'auto' }}
  >
    <Container sx={{ transform: 'translate(0%, -50%)' }}>
      <Paper>{props.children}</Paper>
    </Container>
  </MuiModal>
);
