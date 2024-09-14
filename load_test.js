import http from 'k6/http';
import { sleep } from 'k6';
import { check, group, fail } from 'k6';

export let options = {
    // Configurações de etapas (stages) para simulação de carga progressiva
    stages: [
        { duration: '20s', target: 100 },   // 100 usuários durante 1 minuto
        { duration: '25s', target: 500 },   // 500 usuários durante 2 minutos
        { duration: '30s', target: 2000 },  // 2000 usuários durante 3 minutos
        { duration: '15s', target: 2500 },  // 2500 usuários durante 4 minutos
    ],

    // Limites de tempo para a execução de cada requisição
    thresholds: {
        http_req_duration: ['p(95)<500'], // 95% das requisições devem ter uma duração abaixo de 500ms
        http_req_failed: ['rate<0.01'],   // Taxa de falha das requisições deve ser menor que 1%
    },

    // Tempo máximo de execução do script
    maxDuration: '3m',

    // Configuração de comportamento dos usuários virtuais (VUs)
    vus: 100,  // Número inicial de VUs
    vusMax: 1000,  // Número máximo de VUs possíveis

    // Configurações de retry e fallback para simular erros de rede
    gracefulStop: '3s', // Tempo de espera ao final do teste para finalizar as requisições em andamento
};

export default function () {
    group('Testando endpoint /people', () => {
        const params = {
            headers: {
                "X-API-KEY": "65464",
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVkbm8iLCJleHAiOjE3MjY0Mzg2NTUsImlhdCI6MTcyNjM1MjI1NX0.iLkOQQyUaI0GcjmaGSl3QY7GF6YEr2jFVw5Z2zTZS84"
            }
        }
        const res = http.get('http://localhost:8080/api/v1/people', params);

        // Valida se a resposta é 200 OK
        check(res, {
            'status é 200': (r) => r.status === 200,
        });

        if (res.status !== 200) {
            fail('Falha na requisição');
        }
        sleep(1);
    });
}
