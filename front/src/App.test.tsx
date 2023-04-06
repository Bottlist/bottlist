import { render, screen } from '@testing-library/react';
import { Index } from './pages/Index';
import '@testing-library/jest-dom';

it('should render the title', () => {
  render(<Index />);
  const title = screen.getByText('Bottlist');
  expect(title).toBeInTheDocument();
});
