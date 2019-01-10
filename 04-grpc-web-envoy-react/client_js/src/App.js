import React, { Component } from 'react';
import './App.css';


import {GreeterClient} from './proto/helloworld_grpc_web_pb.js';

import {HelloRequest} from './proto/helloworld_pb.js';

const client = new GreeterClient('http://localhost:9000');

class App extends Component {
  state = {
    response: null
  }
  sayHello = () => {
    const request = new HelloRequest();
    request.setName(this.input.value);

    client.sayHello(request, {}, (err, response) => {
      this.setState({response: response});
    });
  };

  render() {
    const {response} = this.state;
    return (
      <div className="App">
        <header className="App-header">
          <input type="text" ref={el => this.input = el}/>
          <button onClick={this.sayHello}>
            Say Hello
          </button>
          <div>
            Response: { response && response.getMessage() }
          </div>
        </header>
      </div>
    );
  }
}

export default App;
