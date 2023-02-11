import Box from "../../../shared/components/Box";
import { Form } from "../../../shared/components/Form";
import Typography from "../../../shared/components/Typography/indes";
import { useForm, Controller  } from "react-hook-form";
import { FormNotification } from './types';
import { useNotification } from "../services";
import { Textarea } from "../../../shared/components/Textarea";
import { Select } from "../../../shared/components/Select";

function NotificationScreen() {

  const { control, handleSubmit, formState: { errors } } = useForm<FormNotification>({
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

        <Controller
          name="category_id"
          control={control}
          rules={{ required: true }}
          render={({ field }) => (
            <Box marginY="20px">
              <Select
                title="Select the category"                
                {...field}
              />

              {
              errors.text?.message && 
                <Box padding="10px">
                  {errors.text?.message}
                </Box>
              }
            </Box>
          )}
        />

        <Controller
          name="text"
          control={control}
          rules={{ required: true }}
          render={({ field }) => (
            <Box marginY="20px">
              <Textarea
                title="enter the text"
                rows={5}
                {...field}
              />

              {
              errors.text?.message && 
                <Box padding="10px">
                  {errors.text?.message}
                </Box>
              }
            </Box>
          )}
        />

      </Form>
    </Box>
  );
}

export default NotificationScreen;
