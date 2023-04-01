import {Form, Input, Modal} from "antd";
import {EyeInvisibleOutlines, EyeTwoTone} from "@ant-design/icons";
import FormItem from "antd/es/form/FormItem";

const AuthModal = ({shouldShowModal, onSubmit, onCancel}) => {
    const [form] = Form.useForm();

    const onFormSubmit = () => {
        form.validateFields().then((values) => {
            onSubmit(values.token);
        });
    };

    return (
        <Modal
        title ="Provide Github Auth Token"
        centered
        okText="Save"
        cancelText="Cancel"
        open={shouldShowModal}
        onOk={onCancel}
        onCancel={onCancel}
        >

            <Form
                form={form}
                name="auth_form"
                initialValues={{
                    token: "",
                }}
            >
                <Form.Item
                    name="token"
                    label="Token"
                    rules={[{
                        required: true, message:"Please provide your Github Token",
                    },]}
                >
                    <Input.Password 
                        placeholder="Github Token"
                        iconRender={(visible) => visible ? <EyeTwoTone/>}
                    />

                </Form.Item>

            </Form>

        </Modal>
    )

}