import { SSTConfig } from "sst";
import { microservices } from "./stacks/stack";




export default {
  config(_input) {
    return {
      name: "aws-event-microservices",
      region: "us-east-1",
    };
  },

  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });

    
    app.stack(microservices)

  },
} satisfies SSTConfig;
