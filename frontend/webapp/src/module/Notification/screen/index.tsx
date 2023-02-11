import Box from "../../../shared/components/Box";
import { Form } from "../../../shared/components/Form";
import Typography from "../../../shared/components/Typography/indes";
import { useForm } from "react-hook-form";

function NotificationScreen() {

  const { register, handleSubmit, watch, formState: { errors } } = useForm();
  const onSubmit = (data) => console.log(data);

  console.log(watch("example")); // watch input value by passing the name of it


  return (
    <Box
      padding={40}
      backgroundColor="red"
    >
      <Box>
        <Typography>Sending Notifications</Typography>
      </Box>
      <Form onSubmit={handleSubmit(onSubmit)}>      
        <input defaultValue="test" {...register("example")} />
        <input {...register("exampleRequired", { required: true })} />
        {errors.exampleRequired && <span>This field is required</span>}      
        <input type="submit" />    
      </Form>
    </Box>
  );
}

export default NotificationScreen;
