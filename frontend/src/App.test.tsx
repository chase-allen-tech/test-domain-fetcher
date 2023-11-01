// App.test.tsx

import React from 'react';
import { render, act } from '@testing-library/react';
import App from './App';
import axios from 'axios';

// Mocking axios
jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

type Status = {
  url: string;
  statusCode: number;
  duration: number;
  date: string;
};

describe('<App />', () => {
  beforeEach(() => {
    // Resetting mock data
    mockedAxios.get.mockReset();
  });

  it('renders data correctly', async () => {
    const amazonData: Status = {
      url: "amazon-test-url",
      statusCode: 200,
      duration: 100,
      date: "test-date",
    };
    const googleData: Status = {
      url: "google-test-url",
      statusCode: 404,
      duration: 150,
      date: "test-date",
    };
    const allData: Status[] = [amazonData, googleData];

    mockedAxios.get.mockImplementation((url: string) => {
      if (url.includes('amazon-status')) return Promise.resolve({ data: amazonData });
      if (url.includes('google-status')) return Promise.resolve({ data: googleData });
      if (url.includes('all-status')) return Promise.resolve({ data: allData });
      return Promise.reject(new Error('Not found'));
    });

    const { getByText } = render(<App />);

    // Wait for axios requests using act
    await act(async () => {});

    // Check if data is rendered
    expect(getByText("amazon-test-url")).toBeInTheDocument();
    expect(getByText("200")).toBeInTheDocument();
    expect(getByText("100 ms")).toBeInTheDocument();
    expect(getByText("test-date")).toBeInTheDocument();
    expect(getByText("google-test-url")).toBeInTheDocument();
    expect(getByText("404")).toBeInTheDocument();
    expect(getByText("150 ms")).toBeInTheDocument();
  });
});
