import { animate, trigger, query, keyframes, transition, style, state, stagger } from "@angular/animations";
const duration = '0.45s'
const timelineFunction = 'cubic-bezier(.46,.31,.48,.14)'
const space = `${duration} ${timelineFunction}`
const name = 'progress-bar'
const _for = 'true'
const tfor = 'false=>true'
const fromColor = '#C2BD00'
const toColor = '#E5E3E1'

export const PBAnimation = trigger(
    name,
    [
        transition('*=>*', [

            animate(space, keyframes([
                style({ background: fromColor, offset: 0 }),
                style({ background: toColor, offset: 0.5 }),

            ]))

        ])
    ]
)