import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './index.module.css';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import CodeBlock from '@theme/CodeBlock';
import example from '!!raw-loader!./example.go';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <>
          <CodeBlock className={styles.codeblock} language="bash">go get -u github.com/vnteamopen/godebouncer</CodeBlock>
        </>
      </div>
    </header>
  );
}

export default function Home() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title="GoDebouncer">
      <HomepageHeader />
      <main>
        <div className={styles.example}>
          <CodeBlock language="go">{example}</CodeBlock>
        </div>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}