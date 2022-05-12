import React from 'react';
import Link from '@docusaurus/Link';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Run Actions Before Sending Signal',
    description: (
      <>
        Run a special action before the debouncer sends signal with <code>debouncer.Do(specialFunc)</code>. The debouncer first invokes <code>specialFunc</code>, then sends signal to invoke triggered function after <code>wait</code> time.
      </>
    ),
  },
  {
    title: 'Control Debouncer Lifecycle' ,
    description: (
      <>
        Cancel debouncer from invoking the triggered function at any time with <code>debouncer.Cancel()</code>. Send a signal to the debouncer again when you want to restart it.
      </>
    ),
  },
  {
    title: 'Update Debouncer After Sending Signal',
    description: (
      <>
        Debouncer allows replacing the triggered function and the timer after the signal was sent. The new timer take effect in the next <code>SendSignal()</code>.
      </>
    ),
  },
  {
    title: 'Notify the Caller When Triggered Func Finishes',
    description: (
      <>
        Debouncer allows sending a signal via the <code>Done()</code> channel to the caller to let it knows the triggered func has been executed successfully.
      </>
    ),
  },
];

function Feature({title, description}) {
  return (
    <div className={clsx('col col--3')}>
      <div className="text--left padding-horiz--lg">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
        <div className="row">
          <div className={clsx('col col--12')}>
            <div className={styles.buttons}>
              <Link
                className="button button--primary button--lg"
                to="https://github.com/vnteamopen/godebouncer">
                Documentation
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
