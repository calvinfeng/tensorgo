/**
 * @author Calvin Feng
 */

// Libraries
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import axios from 'axios';
import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';

// Components
import BarChart from './components/bar_chart';

// Helpers
import { classifyImageFile } from './util';

// Style
import './index.scss';


interface IndexState {
    currentImage: HTMLImageElement;
    errorMessage: string | undefined;
    x: string[];
    y: number[];
    isLoading: boolean;
}

type ClassResult = {
    prob: number;
    index: number;
    name: string;
}

class Index extends React.Component<any, IndexState> {

    private fileReader: FileReader

    constructor(props) {
        super(props);

        this.state = {
            isLoading: false,
            currentImage: undefined,
            errorMessage: undefined,
            x: [],
            y: []
        };

        this.fileReader = new FileReader();
        this.fileReader.onload = this.handleFileOnLoad;
    }

    handleFileOnLoad = (e: ProgressEvent): void => {
        const img = new Image();
        img.src = this.fileReader.result;

        this.setState({ currentImage: img });
    }

    handleImageFileSelected = (e: React.FormEvent<HTMLInputElement>): void => {
        const imageType = /image.*/;
        const file: File = e.currentTarget.files[0];

        if (file.type.match(imageType)) {
            this.fileReader.readAsDataURL(file);

            classifyImageFile(file).then((results: ClassResult[]) => {
                this.setState({
                    x: results.map((result) => result.name.split(',')[0]),
                    y: results.map((result) => result.prob),
                    isLoading: false
                });
            }).catch((err) => {
                this.setState({
                    errorMessage: err.message
                });
            });
            
            this.setState({ isLoading: true });
        } else {
            this.setState({ errorMessage: "File not supported!" });
        }
    }

    get imageLoader(): JSX.Element {
        let image: JSX.Element;
        if (this.state.currentImage) {
            image = <img src={this.state.currentImage.src} />;
        } else {
            image = <div />;
        }

        return (
            <section className="image-loader">
                <div className="image-container">
                    {image}
                </div>
                <form>
                    <input className="image-input"
                            accept="image/*"
                            id="hidden-image-file-input" 
                            type="file" 
                            name="imageLoader" 
                            onChange={this.handleImageFileSelected} />
                    <label htmlFor="hidden-image-file-input">
                        <Button variant="contained" component="span">Upload</Button>
                    </label>
                </form>
            </section>
        );
    }

    get progresss(): JSX.Element {
        if (this.state.isLoading) {
            return (
                <div className="progress">
                    <CircularProgress />
                </div>
            );
        }

        return <div className="progress"></div>;
    }

    get error(): JSX.Element {
        if (this.state.errorMessage) {
            return <p>{this.state.errorMessage}</p>;
        }

        return <div></div>;
    }

    render() {
        return (
            <section className="index-container">
                <h1>Image Classification</h1>
                <p>Classify an image using residual neural network with 50 layers</p>
                <section className="image-classifier">
                    {this.imageLoader}
                    {this.error}
                    {this.progresss}
                </section>
                <BarChart 
                    title={"Results from ResNet-50"}
                    label={"Softmax Probability"}
                    x={this.state.x}
                    y={this.state.y}  />
            </section>
        )
    }
}


ReactDOM.render(<Index />, document.getElementById('root'));