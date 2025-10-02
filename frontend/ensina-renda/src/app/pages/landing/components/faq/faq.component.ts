import { NgClass, NgFor } from '@angular/common';
import { Component } from '@angular/core';

@Component({
  selector: 'app-faq',
  standalone: true,
  imports: [NgFor, NgClass],
  templateUrl: './faq.component.html',
  styleUrl: './faq.component.css'
})
export class FaqComponent {
  faqs = [
    {
      question: 'O curso é realmente 100% gratuito?',
      answer: 'Sim! Todo o conteúdo do EnsinaRenda é totalmente gratuito. Não existe cobrança futura, planos pagos ou assinaturas escondidas.'
    },
    {
      question: 'Preciso ter conhecimento prévio sobre investimentos para começar?',
      answer: 'Não! O curso foi pensado para iniciantes. Você vai aprender desde o básico com linguagem simples e exemplos práticos.'
    },
    {
      question: 'Já invisto em Renda Fixa. Ainda assim esse curso é útil para mim?',
      answer: 'Com certeza. Mesmo para quem já investe, o curso ajuda a entender melhor os produtos e estratégias de Renda Fixa.'
    },
    {
      question: 'Como funcionam os testes ao final de cada módulo?',
      answer: 'Os testes são personalizados com IA e baseados no conteúdo do módulo, para reforçar o aprendizado.'
    },
    {
      question: 'Preciso fazer os testes para concluir o módulo?',
      answer: 'Sim. Os testes são obrigatórios para garantir que o conteúdo foi absorvido.'
    },
    {
      question: 'Por que preciso criar uma conta para acessar o curso?',
      answer: 'Para salvar seu progresso e permitir acesso contínuo ao conteúdo a qualquer momento.'
    },
    {
      question: 'Em quanto tempo posso concluir o curso?',
      answer: 'Você pode concluir no seu próprio ritmo. O acesso é vitalício.'
    }
  ];

  openedIndex: number | null = null;

  toggleAnswer(index: number) {
    this.openedIndex = this.openedIndex === index ? null : index;
  }

}
