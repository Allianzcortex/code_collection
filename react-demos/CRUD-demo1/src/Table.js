import React, {Component} from 'react';

const TableHeader = () => {
    return (
        <thead>
            <tr>
                <th>name</th>
                <th>Job</th>
                <th>Remove</th>
            </tr>
        </thead>
    );
};

const TableBody = props => {
    const rows = props.characterData.map((row, index) => {
        return (
            <tr key={index}>
                <td>{row.name}</td>
                <td>{row.job}</td>
                <td>
                    <button onClick={() => props.removeCharacter(index)}> delete
                    </button>
                </td>
            </tr>
        );
    });

    return <tbody>{rows}</tbody>;
}

// const TableBody = () => {
//     return (
//         <tbody>
//         {/*<tr>*/}
//         {/*<td>Charlie</td>*/}
//         {/*<td>Janitor</td>*/}
//         {/*</tr>*/}
//         {/*<tr>*/}
//         {/*<td>Mac</td>*/}
//         {/*<td>Bouncer</td>*/}
//         {/*</tr>*/}
//         {/*<tr>*/}
//         {/*<td>Dee</td>*/}
//         {/*<td>Aspiring actress</td>*/}
//         {/*</tr>*/}
//         {/*<tr>*/}
//         {/*<td>Dennis</td>*/}
//         {/*<td>Bartender</td>*/}
//         {/*</tr>*/}
//         </tbody>
//     );
// };

class Table extends Component {

    render() {
        // 通过 Table 传给 TableBody 的所有变量，比如
        // characterData，比如 removeCharacter
        // 都需要通过 props 传递
        const {characterData, removeCharacter} = this.props;
        return (
            <table>
                <TableHeader />
                <TableBody
                    characterData={characterData}
                    removeCharacter={removeCharacter}
                />

            </table>
        );
    }
}

export default Table;