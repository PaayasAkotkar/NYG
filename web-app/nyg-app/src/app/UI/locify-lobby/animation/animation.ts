import { style, transition, trigger, query, group, animate, animateChild, keyframes } from "@angular/animations";
const timelineFunction = 'cubic-bezier(1, .71, 0, -0.06)'
const duration = '0.45s'
export const OnEnter = trigger('OnEnter', [
    transition(':enter', [
        style({ transform: 'translateY(-100%)', }),
        animate(`${duration} ${timelineFunction}`, style({ transform: 'translateY(0)', }))
    ]),
    transition(':leave', [
        animate(`${duration} ${timelineFunction}`, style({ transform: 'translateY(130%)', }))
    ])
]);
