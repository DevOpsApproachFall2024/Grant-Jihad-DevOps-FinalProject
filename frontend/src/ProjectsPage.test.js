import React from 'react';
import { render, screen } from '@testing-library/react';
import ProjectsPage from './ProjectsPage';

test('renders Projects header', () => {
  render(<ProjectsPage />);
  const header = screen.getByText(/Projects/i);
  expect(header).toBeInTheDocument();
});
