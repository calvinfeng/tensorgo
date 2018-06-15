// External imports
import * as React from 'react';
import * as Chart from 'chart.js';

// Internal imports
import { Colors } from './enums';


interface BarChartProps {
    title: string;
    label: string;
    y: number[];
    x: string[];
}


class BarChart extends React.Component<BarChartProps, any> {
    private chart: Chart;

    constructor(props) {
        super(props);
    }

    chartData(props): Chart.ChartData {
        const dataSet = {
            label: props.label,
            backgroundColor: Colors.FETCH_BLUE,
            borderColor: Colors.FETCH_BLUE,
            borderWidth: 1,
            data: props.y
        }

        const chartData: Chart.ChartData = {
            labels: props.x,
            datasets: [dataSet]
        }

        return chartData;
    }

    componentWillReceiveProps(nextProps) {
        this.chart.data = this.chartData(nextProps);
        this.chart.update();
    }

    componentDidMount() {
        const canvas:any = document.getElementById("bar-chart");
        
        const options: Chart.ChartOptions = {
            responsive: true,
            legend: { position: 'top' },
            title: { display: true, text: this.props.title }
        };

        const config: Chart.ChartConfiguration = {
            type: 'bar',
            data: this.chartData(this.props),
            options
        };

        this.chart = new Chart(canvas.getContext('2d'), config);
    }

    render() {
        return <canvas id="bar-chart"></canvas>
    }
}

export default BarChart;