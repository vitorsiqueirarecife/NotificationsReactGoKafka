import Box from "../../../shared/components/Box";
import { Form } from "../../../shared/components/Form";
import Typography from "../../../shared/components/Typography/indes";
import { useForm } from "react-hook-form";
import { FormNotification } from './types';
import { useNotification } from "../services";

function NotificationScreen() {

  const { register, handleSubmit, formState: { errors } } = useForm<FormNotification>({
    defaultValues: {
      category_id: '',
      text: ''
    },
  });

  const { mutate } = useNotification({
    onSuccess: (data) => {
      alert(data.data);
    },
    onError:(e) =>{
      alert(e.message)
    }
  });

  const onSubmit = (data: FormNotification) => {
    mutate(data);
  };

  return (
    <Box
      padding={40}
      backgroundColor="red"
    >
      <Box>
        <Typography>Sending Notifications</Typography>
      </Box>
      <Form onSubmit={handleSubmit(onSubmit)}>              
      </Form>
    </Box>
  );
}

export default NotificationScreen;
