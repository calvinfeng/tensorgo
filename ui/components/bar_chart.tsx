import * as React from 'react';
import * as Chart from 'chart.js';


class BarChart extends React.Component<any, any> {
    private chart: Chart;

    constructor(props) {
        super(props);

        const dataSet = {
            label: 'Data Set 1',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            borderWidth: 1,
            data: [Math.random(), Math.random(), Math.random(), Math.random(), Math.random()]
        };

        const chartData = {
            labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
            datasets: [dataSet]
        };

        this.state = {
            data: chartData
        };
    }

    componentDidMount() {
        const canvas:any = document.getElementById("bar-chart");
        
        const options: Chart.ChartOptions = {
            responsive: true,
            legend: { position: 'top' },
            title: { display: true, text: 'Weekdays' }
        };

        const config: Chart.ChartConfiguration = {
            type: 'bar',
            data: this.state.data,
            options
        };

        this.chart = new Chart(canvas.getContext('2d'), config);
    }

    render() {
        return <canvas id="bar-chart"></canvas>
    }
}

export default BarChart;