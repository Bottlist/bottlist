import { DatePicker as MuiDatePicker } from '@mui/x-date-pickers';
import { Control, Controller, FieldValues, Path } from 'react-hook-form';

export const DatePicker = <T extends FieldValues>(props: {
  _key: Path<T>;
  label: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  control: Control<T, any>;
}) => (
  <Controller
    name={props._key}
    control={props.control}
    render={({ field, fieldState }) => (
      <MuiDatePicker
        label={props.label}
        format="YYYY/MM/DD"
        {...field}
        slotProps={{
          textField: {
            error: !!fieldState.error,
            helperText: fieldState.error?.message,
          },
        }}
      />
    )}
  />
);
