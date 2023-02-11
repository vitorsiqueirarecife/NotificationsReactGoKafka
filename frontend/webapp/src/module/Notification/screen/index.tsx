import Box from "../../../shared/components/Box";
import { Form } from "../../../shared/components/Form";
import Typography from "../../../shared/components/Typography/indes";
import { useForm } from "react-hook-form";
import { FormData } from './types';

function NotificationScreen() {

  const { register, handleSubmit, watch, formState: { errors } } = useForm<FormData>({
    defaultValues: {
      category_id: '',
      text: ''
    },
  });
  const onSubmit = (data: FormData) => console.log(data);

  console.log(watch("category_id"));
  console.log(watch("text"));

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
