import { animate, trigger, query, keyframes, transition, style, state, stagger } from "@angular/animations";
const duration = '0.45s'
const timelineFunction = 'cubic-bezier(.46,.31,.48,.14)'
const space = `${duration} ${timelineFunction}`
const name = 'updrade-power'
const _for = 'true'
const tfor = 'false=>true'
const fromColor = '#C2BD00'
const toColor = '#E5E3E1'

const style1 = { background: `linear-gradient(to right, ${fromColor} {{currentLevel}}%, ${toColor}{{currentLevel}}%`, offset: 1 }
const _class = '.slot'
export const powerAnimation = trigger(
    name,
    [
        transition('*=>*', [
            query(_class,
                [stagger(300,
                    animate(space, keyframes([
                        style({ background: 'red', offset: 0 }),
                        style({ background: 'blue', offset: 0.5 }),

                    ])))]
            )
        ])
    ]
)