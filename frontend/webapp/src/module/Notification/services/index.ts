import { UseMutationOptions } from '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query/build/lib/useMutation';
import { AxiosError } from 'axios';
import apiV1 from '../../../shared/config/axios';
import { FormNotification } from '../screen/types';

import { NotificationResponse } from './types';

export const useNotification = (
  options?: UseMutationOptions<NotificationResponse, AxiosError, FormNotification>,
) => {
  return useMutation<NotificationResponse, AxiosError, FormNotification>(
    (data) => apiV1.post<NotificationResponse>(`/send-message`, data).then((res) => res.data),
    options,
  );
};