import React, {Component} from 'react';
import logo from './logo.svg';
import Table from './Table';
import Form from './Form';

class App extends Component {
    state = {
        characters: []
    };

    removeCharacter = index => {
        const {characters} = this.state;

        this.setState({
            characters: characters.filter((character, i) => {
                return i !== index;
            })
        });
    };

    // characters = [
    //     {
    //         'name': 'Charlie',
    //         'job': 'Janitor'
    //     },
    //     {
    //         'name': 'Mac',
    //         'job': 'Bouncer'
    //     },
    //     {
    //         'name': 'Dee',
    //         'job': 'Aspring actress'
    //     },
    //     {
    //         'name': 'Dennis',
    //         'job': 'Bartender'
    //     }
    // ];

    handleSubmit = character => {
        this.setState({
            characters: [...this.state.characters,
                character]
        });
    }

    render() {
        const {characters} = this.state;
        // return (
        //   <div className="App">
        //     <div className="App-header">
        //         <!-- 注意下面是 className 而不是 class-->
        //         <!-- 换言之是 js 代码而不是 HTML 代码-->
        //       <img src={logo} className="App-logo" alt="logo" />
        //       <h2>Welcome to React</h2>
        //     </div>
        //     <p className="App-intro">
        //       To get started1, edit <code>src/App.js</code> and save to reload.
        //     </p>
        //   </div>
        // );
        return (
            <div className="container">
                <Table
                    characterData={characters}
                    removeCharacter={this.removeCharacter}
                />
                <Form handleSubmit={this.handleSubmit}/>
            </div>
        );
    }
}

// 在 index.js 中会用到
// .render(
//     <App/>,
// 所以在 app.js 里要 export 出
export default App;
