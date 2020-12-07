import { Box, Button, Paper, TextField, Typography } from "@material-ui/core";
import React from "react";
import { useAppContext } from "../AppContext";

export function Settings(): JSX.Element {
  const { endpoint, setEndpoint } = useAppContext();
  const [endpointPending, setEndpointPending] = React.useState<string>(
    endpoint || ""
  );

  return (
    <Box>
      <Paper>
        <Typography variant="h3">Endpoint</Typography>
        <form noValidate>
          <TextField
            label="Endpoint"
            value={endpointPending}
            onChange={(ev) => {
              setEndpointPending(ev.target.value);
            }}
          />
          <Button
            variant="contained"
            onClick={() => {
              if (endpointPending) {
                setEndpoint(endpointPending);
              }
            }}
          >
            Save
          </Button>
        </form>
      </Paper>
    </Box>
  );
}
