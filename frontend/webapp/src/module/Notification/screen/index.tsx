import Box from "../../../shared/components/Box";
import { Form } from "../../../shared/components/Form";
import Typography from "../../../shared/components/Typography/indes";
import { useForm, Controller  } from "react-hook-form";
import { FormNotification } from './types';
import { useNotification } from "../services";
import { Textarea } from "../../../shared/components/Textarea";
import { Select } from "../../../shared/components/Select";
import Button from "../../../shared/components/Button";
import { Option } from "../../../shared/components/Option";

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
      backgroundColor="#F5F5F5"      
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
                {...field}
              >
                <Option>Select the category</Option>
                <Option value="1">Sports</Option>
                <Option value="2">Finance</Option>
                <Option value="3">Movies</Option>
              </Select>

              {
              errors.category_id?.message && 
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
                placeholder="enter the text"                
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

        <Box display="flex" justifyContent="flex-end" width="100%">
          <Button type="submit">Send</Button>
        </Box>
      </Form>
    </Box>
  );
}

export default NotificationScreen;
