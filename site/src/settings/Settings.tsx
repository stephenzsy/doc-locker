import { API, GraphQLResult } from "@aws-amplify/api";
import { Auth, CognitoUser } from "@aws-amplify/auth";
import { Button, Card, Col, Form, Input, Row, Space } from "antd";
import React from "react";
import { useAppContext } from "../AppContext";
import { AuthContext } from "../auth/AuthContext";
import { GetUserProfileQuery } from "../generated/API";
import { getUserProfile } from "../graphql/queries";

export interface IAwsFileDestination {
  name: string;
  cognitoIdentityPoolId: string;
  cognitoIdentityPoolRegion: string;
  s3Bucket: string;
  s3Region: string;
  s3Prefix: string;
}

export interface IAwsSettings {
  version: number;
  destinations: IAwsFileDestination[];
}

interface IAwsFileDestinationProps {
  initialValue: Readonly<IAwsFileDestination> | undefined;
  onChange: (nextValue: IAwsFileDestination) => void;
}

function AwsFileDestinationForm(props: IAwsFileDestinationProps): JSX.Element {
  const { initialValue } = props;
  const [name, setName] = React.useState<string>(initialValue?.name || "");
  const [
    cognitoIdentityPoolId,
    setCognitoIdentityPoolId,
  ] = React.useState<string>(initialValue?.cognitoIdentityPoolId || "");
  const [
    cognitoIdentityPoolRegion,
    setCognitoIdentityPoolRegion,
  ] = React.useState<string>("");
  const [s3BucketName, setS3BucketName] = React.useState<string>("");
  const [s3BucketRegion, setS3BucketRegion] = React.useState<string>("");
  const [s3Prefix, setS3Prefix] = React.useState<string>("");

  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };
  const tailLayout = {
    wrapperCol: { offset: 8, span: 16 },
  };

  const onFinish = React.useCallback((values) => {
    console.log("Success:", values);
  }, []);

  const onFinishFailed = React.useCallback((errorInfo) => {
    console.log("Failed:", errorInfo);
  }, []);

  return (
    <Form<IAwsFileDestination>
      {...layout}
      name="basic"
      initialValues={props.initialValue}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <Form.Item
        label="Name"
        name="name"
        rules={[
          {
            required: true,
            message: "Please input the name of the destination",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}

function AwsSettingsForm(): JSX.Element | null {
  const { aws } = React.useContext(AuthContext);
  const { endpoint } = useAppContext();
  const [currentUser, setCurrentUser] = React.useState<CognitoUser | undefined>(
    undefined
  );
  React.useEffect(() => {
    if (aws) {
      (async () => {
        const user = (await Auth.currentAuthenticatedUser()) as CognitoUser;
        if (user) {
          setCurrentUser(user);
        }
        const result = (await API.graphql({
          query: getUserProfile,
          variables: { id: "me" },
        })) as GraphQLResult<GetUserProfileQuery>;
        const configsString = result.data?.getUserProfile?.awsConfigs;
        if (!configsString) {
          return;
        }
        const awsSettings = JSON.parse(configsString) as IAwsSettings;
        if (awsSettings.version === 1) {
          // load to UI, otherwise non compatible
        }
      })();
    }
  }, [aws, endpoint]);

  if (!aws) {
    return null;
  }

  return (
    <AwsFileDestinationForm
      initialValue={undefined}
      onChange={() => {
        // do nothing
      }}
    />
  );
}

export function Settings(): JSX.Element {
  const { endpoint, setEndpoint } = useAppContext();
  const [endpointPending, setEndpointPending] = React.useState<string>(
    endpoint || ""
  );

  return (
    <Row>
      <Col span={24}>
        <Space>
          <Card title="Endpoint">
            <Form>
              <Form.Item label="Endpoint">
                <Input
                  value={endpointPending}
                  onChange={(ev) => {
                    setEndpointPending(ev.target.value);
                  }}
                />
              </Form.Item>
              <Form.Item>
                <Button
                  onClick={() => {
                    if (endpointPending) {
                      setEndpoint(endpointPending);
                    }
                  }}
                >
                  Save
                </Button>
              </Form.Item>
            </Form>
          </Card>
          <Card title="AWS Settings">
            <AwsSettingsForm />
          </Card>
        </Space>
      </Col>
    </Row>
  );
}
