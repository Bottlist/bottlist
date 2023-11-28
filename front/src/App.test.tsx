import { render, screen } from '@testing-library/react';
import { Index } from './pages/Index';
import { BrowserRouter } from 'react-router-dom';
import '@testing-library/jest-dom';

it('should render the title', () => {
  render(
    <BrowserRouter>
      <Index />
    </BrowserRouter>
  );
  const title = screen.getByText('BOTTLIS');
  expect(title).toBeInTheDocument();
});
