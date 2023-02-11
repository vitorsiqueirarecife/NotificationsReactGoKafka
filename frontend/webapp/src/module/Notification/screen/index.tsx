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

  const { mutate, isLoading } = useNotification({    
    onSuccess: (data) => {
      alert(data.message);
    },
    onError:(e) =>{
      alert(e.message)
    }
  });

  const onSubmit = (data: FormNotification) => {    
    if(!isLoading){
      mutate(data);
    }
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
          rules={{ min: 1, max: 3, required: true }}
          render={({ field }) => (
            <Box marginY="20px">
              <Select             
                borderColor={errors?.category_id ? 'red' : undefined}        
                {...field}
              >
                <Option value="0">Select the category</Option>
                <Option value="1">Sports</Option>
                <Option value="2">Finance</Option>
                <Option value="3">Movies</Option>
              </Select>             
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
                borderColor={errors?.text ? 'red' : undefined}
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
          <Button type="submit" >Send</Button>
        </Box>
      </Form>
    </Box>
  );
}

export default NotificationScreen;
