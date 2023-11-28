import {
  TableContainer as MuiTableContainer,
  TableCell,
  TableCellProps,
} from '@mui/material';

export const TableContainer = (props: { children: JSX.Element }) => (
  <MuiTableContainer
    sx={{
      borderTop: '1px solid #555555',
      borderBottom: '1px solid #555555',
    }}
  >
    {props.children}
  </MuiTableContainer>
);

export const Td = (props: TableCellProps) => (
  <TableCell
    sx={{
      borderTop: '1px inset solid #555555',
      borderBottom: '1px inset solid #555555',
    }}
    {...props}
  />
);
export const Th = (props: TableCellProps) => (
  <TableCell {...props} component="th" scope="row" sx={{ border: 'none' }} />
);
